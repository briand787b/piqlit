resource "aws_iam_role" "codebuild_role" {
    name = "codebuild_role"
    assume_role_policy = <<-EOF
    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {
                    "Service": "codebuild.amazonaws.com"
                },
                "Action": "sts:AssumeRole"
            }
        ]
    }
    EOF
}

resource "aws_iam_role_policy" "codebuild_role_policy" {
    name = "codebuild_role_policy"
    role = aws_iam_role.codebuild_role.id

    policy = <<-EOF
    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Resource": [
                    "arn:aws:logs:us-east-1:565527435302:log-group:/aws/codebuild/piqlit",
                    "arn:aws:logs:us-east-1:565527435302:log-group:/aws/codebuild/piqlit:*"
                ],
                "Action": [
                    "logs:*"
                ]
            },
            {
                "Effect": "Allow",
                "Resource": [
                    "arn:aws:s3:::codepipeline-us-east-1-*"
                ],
                "Action": [
                    "s3:PutObject",
                    "s3:GetObject",
                    "s3:GetObjectVersion",
                    "s3:GetBucketAcl",
                    "s3:GetBucketLocation"
                ]
            }
        ]
    }
    EOF
}