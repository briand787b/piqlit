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
                    "${var.codebuild_log_group.arn}"
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