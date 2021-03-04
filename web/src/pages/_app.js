import { ChakraProvider, ColorModeProvider } from "@chakra-ui/react";
import { Provider } from "react-redux";
import Wrapper from "../components/Wrapper";
import { useStore } from "../redux/store";
import theme from "../theme";
function MyApp({ Component, pageProps }) {
  const store = useStore();
  return (
    <Provider store={store}>
      <ChakraProvider resetCSS theme={theme}>
        <ColorModeProvider
          options={{
            useSystemColorMode: true,
          }}
        >
          <Wrapper>
            <FullHeightPage>
              <Component {...pageProps} />
            </FullHeightPage>
          </Wrapper>
        </ColorModeProvider>
      </ChakraProvider>
    </Provider>
  );
}
const FullHeightPage = ({ children }) => (
  <div>
    {children}
    <style global jsx>{`
      html,
      body,
      body > div:first-child,
      div#__next,
      div#__next > div {
        height: 95%;
      }
    `}</style>
  </div>
);
export default MyApp;
