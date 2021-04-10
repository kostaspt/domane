import React from 'react';
import { render as rtlRender } from '@testing-library/react';
import { configureStore } from '@reduxjs/toolkit';
import { Provider } from 'react-redux';
import reducer from '@store/reducer';

const makeStore = (initialState: any) => configureStore({ reducer, preloadedState: initialState });

function renderWithStore(
  component: React.ReactElement,
  { initialState, store = makeStore(initialState), ...renderOptions } = {} as any
) {
  function Wrapper({ children }: any) {
    return <Provider store={store}>{children}</Provider>;
  }

  return rtlRender(component, { wrapper: Wrapper, ...renderOptions });
}

export * from '@testing-library/react';
export { renderWithStore };
