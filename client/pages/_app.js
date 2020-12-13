import { Grommet } from 'grommet';
import getConfig from 'next/config';
import { Provider } from 'react-redux';
import store from '../store';

const { publicRuntimeConfig } = getConfig();
const { THEME } = publicRuntimeConfig;

export default function CustomApp({ Component, pageProps }) {
  return (
    <Grommet theme={THEME}>
      <Provider store={store}>
        <Component {...pageProps} />
      </Provider>
    </Grommet>
  );
}
