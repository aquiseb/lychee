MAKEFLAGS += --silent

define GetVal
	$(shell node -p '\
		const getVal = (key = "", obj = {}) => {
			const currKey = key.split(".")[0];
			const val = obj[currKey];
			if(typeof val !== "object") return val;
			const nextKey = key.split(".").slice(1).join(".");
			return getVal(nextKey, val);
		}; \
		getVal(`$(1)`.replace(" ", ""), require("./package.json")); \
	' --no-warnings)
endef

APP_NAME := $(call GetVal, name)
DEPLOYMENT_NAME= $(subst @aquiseb/,,$(APP_NAME))

show:
	echo $(DEPLOYMENT_NAME)

hero-env-get:
	-heroku config -s -a $(DEPLOYMENT_NAME) > .env.heroku
	
hero-env-reset:
	-cat .env.heroku | cut -d'#' -f1 | tr '=' ' ' | xargs heroku config:unset -a $(DEPLOYMENT_NAME) > /dev/null 2>&1
	rm -rf .env.heroku

hero-env: hero-env-get hero-env-reset
	-cat .env | cut -d'#' -f1 | tr '\n' ' ' | xargs heroku config:set -a $(DEPLOYMENT_NAME) > /dev/null 2>&1

hero-buildpacks:
	-heroku buildpacks:add heroku/nodejs -a $(DEPLOYMENT_NAME)
	-heroku buildpacks:add https://github.com/uhray/heroku-buildpack-mongo.git -a $(DEPLOYMENT_NAME)

hero-deploy: show hero-env
	heroku git:remote -a $(DEPLOYMENT_NAME)
	heroku config:set DEPLOYMENT_ENV=heroku X_HOST=LIVE USE_NPM_INSTALL=true
	git push heroku main --force
