import { Grommet } from 'grommet';
import type { AppProps } from 'next/app';
import getConfig from 'next/config';
import { Provider } from 'react-redux';
import store from '../store';

const { publicRuntimeConfig } = getConfig();
const { THEME } = publicRuntimeConfig;

export default function CustomApp({ Component, pageProps }: AppProps) {
  return (
    <Grommet theme={THEME}>
      <Provider store={store}>
        <Component {...pageProps} />
      </Provider>
    </Grommet>
  );
}
