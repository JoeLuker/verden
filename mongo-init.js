db.createUser({
    user: _getEnv('MONGO_APP_USERNAME'),
    pwd: _getEnv('MONGO_APP_PASSWORD'),
    roles: [
      {
        role: "readWrite",
        db: "_getEnv('MONGO_INITDB_DATABASE')"
      }
    ]
  });
  
  // Utility function to access environment variables
  function _getEnv(varName) {
    return db.getSiblingDB('admin').runCommand({
      $eval: `process.env['${varName}']`
    }).retval;
  }