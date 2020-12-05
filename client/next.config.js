const { PHASE_PRODUCTION_BUILD } = require('next/constants');

module.exports = (phase) => {
  const isProd = phase === PHASE_PRODUCTION_BUILD;

  return {
    publicRuntimeConfig: {
      API_URL: isProd ? '/api' : 'http://localhost:4000',
      THEME: {
        global: {
          colors: {
            brand: '#047857',
          },
          font: {
            family: `ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";`,
          },
        },
      },
    },
  };
};
