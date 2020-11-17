## ustream-heroku

just like https://github.com/suconghou/ustream

this is ustream for heroku


### deploy your own ustream

`git clone https://github.com/suconghou/u2web`

modify `src/service/index.js` , replace `abs` and `apiBaseURL`'s domain to your heroku domain

run `yarn` install dependencies & `yarn build` to build

copy all files in `dist` to `public` dir

you may add `YOUTUBE_API_KEY` env variable to use youtube data api

if you use p2p feature,  you may build `fastloadjs` from source to use your own signalserver and turnserver , update `index.html` `play.html` then rebuild u2web



