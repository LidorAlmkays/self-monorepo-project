import env from './.env';

export const environment = {
  production: false,
  apiBaseUrl: 'http://localhost:8000',
  version: env.npm_package_version,
};
