import { ChakraProvider, ColorModeProvider } from '@chakra-ui/react';
import type { AppProps } from 'next/app';
import { Provider } from 'react-redux';
import store from '../store';
import theme from '../theme';

export default function CustomApp({ Component, pageProps }: AppProps) {
  return (
    <ChakraProvider resetCSS theme={theme}>
      <ColorModeProvider
        options={{
          useSystemColorMode: true,
        }}
      >
        <Provider store={store}>
          <Component {...pageProps} />
        </Provider>
      </ColorModeProvider>
    </ChakraProvider>
  );
}
