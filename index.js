const core = require('@actions/core');

const setup = require('./lib/setup-go-action');

(async () => {
  try {
    await setup();
  } catch (error) {
    core.setFailed(error.message);
  }
})();