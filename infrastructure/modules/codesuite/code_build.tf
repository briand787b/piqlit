resource "aws_s3_bucket" "build_cache" {
    bucket = "piqlit-codebuild-cache"
}

resource "aws_cloudwatch_log_stream" "codebuild_log_stream" {
    name = "codebuild_log_stream"
    log_group_name = var.codebuild_log_group.name
} 

resource "aws_codebuild_project" "codebuild" {
    name = "piqlit-builder"
    description = "builds and tests all containers for piqlit"
    service_role = aws_iam_role.codebuild_role.arn

    source {
        type = "GITHUB"
        location = data.github_repository.piqlit.http_clone_url
    }

    artifacts {
        type = "NO_ARTIFACTS"
    }

    environment {
        type = "LINUX_CONTAINER"
        image = "aws/codebuild/amazonlinux2-x86_64-standard:1.0"
        compute_type = "BUILD_GENERAL1_SMALL"
        
        environment_variable {
            name = "POSTMAN_API_KEY"
            value = var.postman_api_key
        }   

        environment_variable {
            name = "POSTMAN_COLLECTION_ID"
            value = var.postman_collection_id
        }
    }

    cache {
        type = "S3"
        location = aws_s3_bucket.build_cache.arn
    }

    logs_config {
        cloudwatch_logs {
            group_name = var.codebuild_log_group.name
            stream_name = aws_cloudwatch_log_stream.codebuild_log_stream.name
        }
    }
}

resource "aws_codebuild_webhook" "piqlit" {
    project_name = aws_codebuild_project.codebuild.name

    filter_group {
        filter {
            type = "EVENT"
            pattern = "PUSH"
        }
    }
}