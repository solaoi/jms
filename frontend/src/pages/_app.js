import { ApolloProvider, ApolloClient, InMemoryCache } from "@apollo/client";

// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
function MyApp({ Component, pageProps }) {
  const cache = new InMemoryCache();
  const client = new ApolloClient({
    uri: "http://localhost:3000/query",
    cache,
  });

  return (
    <ApolloProvider client={client}>
      <Component {...pageProps} />
    </ApolloProvider>
  );
}

export default MyApp;
