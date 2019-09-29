# Minimal Next + Relay

Install it:

```bash
npm install
# or
yarn
```

Run the project (it runs automatically the Relay ahead-of-time compilation)

```bash
npm run dev
# or
yarn dev
```

## Example queries

Visit `localhost:3000/graphql` and use the following queries

```
{
    allPosts(first: 2) {
        edges {
            node {
                title
                content
            }
        }
    }
}

```

```
{
  viewer {
        isLoggedIn
        user {
            firstname
            lastname
            posts {
                edges {
                    node {
                        title
                        content
                    }
                }
            }
        }
    }
}
```

# React-Native-Web sources
[Build Mobile-Friendly Web Apps with React Native Web](https://scotch.io/tutorials/build-mobile-friendly-web-apps-with-react-native-web)
[From Zero to Publish: Expo Web (React Native for Web) Tutorial](https://medium.com/@toastui/from-zero-to-publish-expo-web-react-native-for-web-tutorial-e3e020d6d3ff)
[Video: React Native Web Full App Tutorial - Build a Workout App for iOS, Android, and Web](https://www.youtube.com/watch?v=_CBYbEGvxYY)
[react-navigation vs react-native-navigation](https://blog.logrocket.com/react-navigation-vs-react-native-navigation-which-is-right-for-you-3d47c1cd1d63/)
[Being free from expo in react-native apps](https://medium.com/reactbrasil/being-free-from-expo-in-react-native-apps-310034a3729)
[Install java on Ubuntu](https://thishosting.rocks/install-java-ubuntu/)

## Sources
https://medium.com/front-end-weekly/learn-using-jwt-with-passport-authentication-9761539c4314
https://stackoverflow.com/questions/39163413/node-js-passport-jwt-how-to-send-token-in-a-cookie
!! https://stackoverflow.com/a/39824342/9077800
https://stackoverflow.com/questions/8455272/expressjs-secure-session-cookie
https://dev.to/jolvera/user-authentication-with-nextjs-4023
https://github.com/zeit/next.js/blob/a2b2a2a3e85808f408f7a507cbce4ceddc64f0ea/examples/with-cookie-auth/utils/auth.js
https://expressjs.com/en/advanced/best-practice-security.html


https://medium.com/the-guild/authentication-and-authorization-in-graphql-and-how-graphql-modules-can-help-fadc1ee5b0c2
https://www.freecodecamp.org/news/how-to-nail-social-authentication-in-graphql-27943aee5dce/
https://www.youtube.com/watch?v=qjKZYQih288
https://dev.to/eveporcello/github-authorization-with-graphql-and-apollo-server-3hf7
https://developers.facebook.com/docs/facebook-login/manually-build-a-login-flow/
https://medium.com/@evangow/server-authentication-basics-express-sessions-passport-and-curl-359b7456003d
https://github.com/rupalipemare/Todo-Demo/search?q=session&unscoped_q=session

## Why graphql
https://www.prisma.io/blog/top-5-reasons-to-use-graphql-b60cfa683511
https://medium.com/@JeffLombardJr/when-and-why-to-use-graphql-24f6bce4839d


## Todo

Error: automatic static optimization failed: found page without a React Component as default export in
pages/**generated**/pages_indexQuery.graphql

See https://err.sh/zeit/next.js/page-without-valid-component for more info.

    at build (/home/mies/minimal-next-relay/node_modules/next/dist/build/index.js:7:3807)
    at process._tickCallback (internal/process/next_tick.js:68:7)

error Command failed with exit code 1.
