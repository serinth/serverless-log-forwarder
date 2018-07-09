# Serverless Log Forwarder Example

This creates a lambda function and a special log group that other serverless functions will make a reference to so that their logs get aggregated in one place. 

The sample forwarder only outputs the log but that's where we would create an integration function into something like Splunk, Logz.io etc.

This project should be deployed first so that the log group ARN is exposed for other serverless functions.

Provider: AWS

# Deployment

With MFA enabled on the AWS account, we need to first grab a temporary session token and use it. 

1. Get an AWS Temporary Session Token:

```bash
    aws sts get-session-token --serial-number <MFA ARN> --token-code <MFA AUTH CODE>
```

2. Update AWS Credentials and Profile files:

Once you have the output, put it in your credentials file and ensure the named profile file also has the MFA ARN listed.

Example `~/.aws/credentials`:
```
    [Default]
    aws_access_key_id = xxx
    aws_secret_access_key = xxx
    
    [TEMPSESSION]
    aws_access_key_id = xxx
    aws_secret_access_key = xxx
    aws_session_token = xxx
```

Example `~/.aws/config`:
```
    [profile TEMPSESSION]
    region = ap-southeast-2
    source_profile = TEMPSESSION
    mfa_serial = <ARN>
```

3. Run Serverless Deployment

```bash
    serverless deploy -v --aws-profile MyProfile
```
# Integration into other Serverless Project

After the logging function has been deployed, we need to take the ARN of the log group from the output and plug it into our other serverless project. Additionally, we need to use the serverless-log-forwarding plugin:

1. `npm install --save-dev serverless-log-forwarding`

2. Modify serverless.yml to have:
```yaml
plugins:
  - serverless-log-forwarding

custom:
  logForwarding:
    destinationARN: <forwarding-function-ARN>
    filterPattern: "-\"RequestId: \""
```

# Clean Up

```bash
    serverless remove -v --aws-profile TEMPSESSION
```