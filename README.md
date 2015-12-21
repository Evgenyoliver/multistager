# Multistager
Multistaging orchestrator

### Start new branch

`$ http POST multistager.service.consul:8080/v1/container image=qlean-staging key=<GITHUB_KEY> branch=<GIT_BRANCH>`

### Restart already deployed branch

`$ http PUT multistager.service.consul:8080/v1/container image=qlean-staging key=<GITHUB_KEY> branch=<GITHUB_BRANCH>`

### Stop branch

`$ http DELETE multistager.service.consul:8080/v1/container branch=<GITHUB_BRANCH>`

### List branches

`$ http GET multistager.service.consul:8080/v1/container image=qlean-staging`
