# mela
## Create Table
aws dynamodb create-table --region us-east-1 --table-name content --attribute-definitions AttributeName=id,AttributeType=S --key-schema AttributeName="id",KeyType="HASH" --billing-mode PAY_PER_REQUEST --endpoint-url https://dynamodb.us-east-1.amazonaws.com
aws dynamodb create-table --region us-east-1 --table-name user --attribute-definitions AttributeName=id,AttributeType=S --key-schema AttributeName="id",KeyType="HASH" --billing-mode PAY_PER_REQUEST --endpoint-url https://dynamodb.us-east-1.amazonaws.com
aws dynamodb create-table --region us-east-1 --table-name spending --attribute-definitions AttributeName=content_id,AttributeType=S AttributeName=user_id,AttributeType=S  --key-schema AttributeName="content_id",KeyType="HASH" AttributeName="user_id",KeyType="RANGE" --billing-mode PAY_PER_REQUEST --endpoint-url https://dynamodb.us-east-1.amazonaws.com

# Delete Table
aws dynamodb delete-table --region us-east-1 --table-name content --endpoint-url https://dynamodb.us-east-1.amazonaws.com
aws dynamodb delete-table --region us-east-1 --table-name user --endpoint-url https://dynamodb.us-east-1.amazonaws.com
aws dynamodb delete-table --region us-east-1 --table-name spending --endpoint-url https://dynamodb.us-east-1.amazonaws.com