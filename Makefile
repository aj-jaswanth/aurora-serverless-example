tf-apply: lambda-build
	cd terraform; terragrunt apply -auto-approve -parallelism 50

tf-destroy:
	cd terraform; terragrunt destroy -auto-approve -parallelism 50

lambda-build:
	cd hello-lambda; GOOS=linux go build; rm function.zip; zip function.zip migrations/* hello-lambda