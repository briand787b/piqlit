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
                    "${aws_codebuild_project.codebuild.arn}",
                    "${aws_codebuild_project.codebuild.arn}:*",
                    "arn:aws:codebuild:us-east-1:565527435302:build/piqlit-builder:*",
                    "arn:aws:codebuild:us-east-1:565527435302:build/piqlit-builder:a60ac2d0-8175-4960-9a57-368b4a39a2a7"
                ],
                "Action": [
                    "logs:CreateLogGroup",
                    "logs:CreateLogStream",
                    "logs:PutLogEvents"
                ]
            }
        ]
    }
    EOF
}