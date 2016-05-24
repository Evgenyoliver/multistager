# Multistager
Multistaging orchestrator

## Usage

### List of running containers

`$ curl -s multistager.service.consul/v1/container?image=qlean-staging | python -m json.tool`

### Start new branch

`$ curl -X POST multistager.service.consul/v1/container -d '{"image":"qlean-staging", "key":"<GITHUB_KEY>", "branch":"<GIT_BRANCH>"}'`

### Restart already deployed branch

`$ curl -X PUT multistager.service.consul/v1/container -d '{"image":"qlean-staging", "key":"<GITHUB_KEY>", "branch":"<GIT_BRANCH>"}'`

### Stop branch

`$ curl -X DELETE multistager.service.consul/v1/container -d '{"image":"qlean-staging", "key":"<GITHUB_KEY>", "branch":"<GIT_BRANCH>"}'`

## Configuration

### Stages limit

`$ curl localhost:8500/v1/kv/multistager/stages_limit?raw`

### Linked containers (splitted by comma)

*For example: messenger-redis:redis, postgres:postgres*

`$ curl localhost:8500/v1/kv/multistager/links?raw`

### Mount folder to container

`curl -X PUT multistager.service.consul/v1/container -d
'{
    "image":"image-name",
    ...
    "flags": ["MountFolder"],
    "mount_folder_path": "/dumps"
}'`

Bind /dumps folder on host machine with multistager to /dumps in container
