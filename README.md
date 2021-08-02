# How To Run ?

Be Sure That `docker` and `docker-compose` are installed  
Just Run `docker-compose up -d --build` In The Root Of Project and visit `0.0.0.0:6699/index` or `0.0.0.0:6699/admin`

## How To Be Sure Load Balancing Working Correctly ?

Each Instance Of Application Has UUID, If You See Different UUIDs In Response, Then Load Balancing Working Perfectly
