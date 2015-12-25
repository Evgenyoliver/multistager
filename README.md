# Multistager
Multistaging orchestrator

### List of running containers

`$ curl -s multistager.service.consul/v1/container?image=qlean-staging | python -m json.tool'`

### Start new branch

`$ curl -X POST multistager.service.consul/v1/container -d '{"image":"qlean-staging", "key":"<GITHUB_KEY>", "branch":"<GIT_BRANCH>"}'`

### Restart already deployed branch

`$ curl -X PUT multistager.service.consul/v1/container -d '{"image":"qlean-staging", "key":"<GITHUB_KEY>", "branch":"<GIT_BRANCH>"}'`

### Stop branch

`$ curl -X DELETE multistager.service.consul/v1/container -d '{"image":"qlean-staging", "key":"<GITHUB_KEY>", "branch":"<GIT_BRANCH>"}'`
