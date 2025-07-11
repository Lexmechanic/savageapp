# savageapp
Savageapp is a small project to implement a local server for playing the Savage Worlds RPG. No actual Savage World IP will ever be committed to this repo this only creates the framework for entering stat blocks of and components of the game system as well as providing a round based combat simulator.

## Structure
Frontend - Svelte 5 running on port 3000
Logic - Go handles APIs/Websocket calls from Frontend runing on port 8080
DB - Sqlite
Network - Recommend using nginx proxy with the nginx_reverse_proxy.config to allow ingress on port 80 and resolve CORS issues

## Running
1. go run . 
2. frontend/npm run dev
3. Configure incoming for nginx if 80 is not desired in nginx_reverse_proxy.config and add to nginx config settings on system

## TODO
- Generate make files for build production executable versions and add docker deployment