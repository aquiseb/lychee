require("dotenv").config();
const withCss = require("@zeit/next-css");

module.exports = withCss({
  serverRuntimeConfig: {
    SERVER_SECRET: process.env.SERVER_SECRET,
    auth: {
      FACEBOOK_CLIENTID: process.env.FACEBOOK_CLIENTID,
      FACEBOOK_CLIENTSECRET: process.env.FACEBOOK_CLIENTSECRET
    },
    app: {
      URL: process.env.NGROK || process.env.APP_URL
    }
  },
  webpack: (config, { isServer }) => {
    if (isServer) {
      const antStyles = /antd\/.*?\/style\/css.*?/;
      const origExternals = [...config.externals];
      config.externals = [
        (context, request, callback) => {
          if (request.match(antStyles)) return callback();
          if (typeof origExternals[0] === "function") {
            origExternals[0](context, request, callback);
          } else {
            callback();
          }
        },
        ...(typeof origExternals[0] === "function" ? [] : origExternals)
      ];

      config.module.rules.unshift({
        test: antStyles,
        use: "null-loader"
      });
    }
    return config;
  }
});
