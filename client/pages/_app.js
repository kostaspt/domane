import { Grommet, ResponsiveContext } from 'grommet';
import { Provider } from 'react-redux';
import store from '../store';

const theme = {
  global: {
    colors: {
      brand: '#10B981',
    },
    font: {
      family: `ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";`,
    },
  },
};

export default function Domane({ Component, pageProps }) {
  return (
    <Grommet theme={theme}>
      <ResponsiveContext.Consumer>
        {(size) => (
          <Provider store={store}>
            <Component size={size} {...pageProps} />
          </Provider>
        )}
      </ResponsiveContext.Consumer>
    </Grommet>
  );
}
