import { ChakraProvider, cookieStorageManager, localStorageManager } from '@chakra-ui/react';
import App, { AppContext } from 'next/app';
import type { AppProps } from 'next/app';
import { Provider } from 'react-redux';
import store from '../store';
import theme from '../theme';

type CustomAppProps = AppProps & {
  cookies: any;
};

function CustomApp({ Component, cookies, pageProps }: CustomAppProps) {
  return (
    <ChakraProvider
      resetCSS
      theme={theme}
      colorModeManager={typeof cookies === 'string' ? cookieStorageManager(cookies) : localStorageManager}
    >
      <Provider store={store}>
        <Component {...pageProps} />
      </Provider>
    </ChakraProvider>
  );
}

CustomApp.getInitialProps = async (appContext: AppContext) => {
  const appProps = await App.getInitialProps(appContext);

  return { ...appProps, cookies: appContext.ctx.req?.headers.cookie };
};

export default CustomApp;
