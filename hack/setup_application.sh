#!/bin/bash

# PRECONDITIONS
#
# It is assumed that:
# 1. that you have setup a personal github access token (https://github.com/settings/tokens)
#    with admin:repo_hook scope and set up the environment
#    variables:
#    GH_USERNAME=<github user name>
#    GITHUB_PAT_TOKEN=<token generated>

REPOSITORY_PATH="<owner>/<repository>"
REPOSITORY="https://github.com/$REPOSITORY_PATH"
OWNER="<email of owner>"
SHARED_SECRET=<Any shared secret to be used by the webhook>
WEBHOOK_HOSTNAME=webhook.radix.equinor.com|webhook.playground.radix.equinor.com|webhook.us.radix.equinor.com

# Set access token for use in call to command line
# Users should get this as part of their Radix application
export APP_SERVICE_ACCOUNT_TOKEN=

echo "Create application"
PUBLIC_KEY=$(rx create application \
    --token-environment \
    --from-config \
    --repository $REPOSITORY \
    --owner $OWNER \
    --shared-secret $SHARED_SECRET 2>&1)

# Wait for application to be reconciled
sleep 3

echo "Add webhook"
RESPONSE=$(curl -X POST -H "Content-Type: application/json" -u ${GH_USERNAME}:${GITHUB_PAT_TOKEN} \
    https://api.github.com/repos/$REPOSITORY_PATH/hooks \
    -d '{"name":"web", "active": true, "config": { "url": "https://'${WEBHOOK_HOSTNAME}/events/github'", "content_type": "json", "secret": "'$SHARED_SECRET'" }}' 2>&1)

echo "Add deploy key"
PAYLOAD='{"title":"playground", "key": "'$PUBLIC_KEY'" }'
RESPONSE=$(curl -X POST -H "Content-Type: application/json" -u ${GH_USERNAME}:${GITHUB_PAT_TOKEN} \
    https://api.github.com/repos/$REPOSITORY_PATH/keys \
    -d "$PAYLOAD" 2>&1)

echo "Build from master branch"
rx build-deploy \
    --token-environment \
    --from-config \
    -f

sleep 3
echo "Set environment secret"
rx set environment-secret \
    --token-environment \
    --from-config \
    --await-reconcile \
    -e <your environment> \
    --component <your component> \
    -s <your secret name> \
    -v '<your secret value>'

echo ""
read -p "Delete application again? (Y/n) " -n 1 -r
if [[ "$REPLY" =~ (N|n) ]]; then
    echo ""
    echo "Chicken!"
fi

rx delete application \
    --from-config
