const { Client } = require('pg');
const { v4: uuidv4 } = require('uuid');

const client = new Client({
    host: 'localhost',
    user: 'postgres',
    database: 'postgres',
    password: 'postgres',
    port: 5432,
});

var dbConn = null;

const getConnection = async () =>  {
    if ( dbConn ==null){
        await client.connect(); 

        dbConn = client;
    }

    return dbConn;
}

const insertUser = async (request) => {
    var client = await getConnection();
    console.log(request);
  
    try {
                  // gets connection
        await client.query(
            `insert into user_login_reports(id, user_pool_id, cognito_user_id, region, email) 
            VALUES ($1, $2, $3, $4, $5)`,	[uuidv4(), request.userPoolID, request.userName, request.region, request.request.userAttributes.email]); // sends queries
        return true;
    } catch (error) {
        console.error(error.stack);
        return false;
    }
};

exports.handler =  async function(event, context,callback) {

    // Send post authentication data to Cloudwatch logs
    console.log ("Authentication successful");
    console.log ("Trigger function =", event.triggerSource);
    console.log ("User pool = ", event.userPoolId);
    console.log ("App client ID = ", event.callerContext.clientId);
    console.log ("User ID = ", event.userName);

    console.log("status", insertUser(event));

    // Return to Amazon Cognito
    callback(null, event);
}