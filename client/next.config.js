const { PHASE_PRODUCTION_BUILD } = require('next/constants');

module.exports = (phase) => {
  const isProd = process.env.NODE_ENV === 'production' || phase === PHASE_PRODUCTION_BUILD;

  const apiDomain = process.env.DOMAIN ?? 'domane.io';
  const apiUrl = apiDomain.match(/\.local(host)?$/) ? `http://${apiDomain}` : `https://${apiDomain}`;

  return {
    publicRuntimeConfig: {
      API_URL: isProd ? `${apiUrl}/api` : 'http://localhost:4001',
    },
    webpack(config) {
      config.module.rules.push({
        test: /\.svg$/,
        use: ['@svgr/webpack'],
      });

      return config;
    },
  };
};
