using Amazon.Lambda.Core;
using Npgsql;

// Assembly attribute to enable the Lambda function's JSON input to be converted into a .NET class.
[assembly: LambdaSerializer(typeof(Amazon.Lambda.Serialization.SystemTextJson.DefaultLambdaJsonSerializer))]

namespace lambda_dotnet;

public class Function
{
    
    String connString = "Host=localhost;Port=5452;Username=postgres;Password=postgres;Database=postgres";

    public async Task FunctionHandlerAsync(ILambdaContext context)
    {

        var conn = await GetConnection();

        // Insert some data
        await using (var cmd = new NpgsqlCommand("insert into user_login_reports(id, user_pool_id, cognito_user_id, region, email, user_attributes)  VALUES ($1, $2, $3, $4, $5, $6)", conn))
        {
            cmd.Parameters.AddWithValue(Guid.NewGuid().ToString());
            cmd.Parameters.AddWithValue("poolId-dotnet");
            cmd.Parameters.AddWithValue("username");
            cmd.Parameters.AddWithValue("eu-west-2");
            cmd.Parameters.AddWithValue("r@ayyildiz.ai");
            cmd.Parameters.AddWithValue("attrs");

            await cmd.ExecuteNonQueryAsync();
        }
    }

    private NpgsqlConnection? dbConn=null;

    private async Task<NpgsqlConnection> GetConnection() {

        if ( dbConn ==null) {
                var conn = new NpgsqlConnection(connString);
                await conn.OpenAsync();
                dbConn = conn;

        }

        return dbConn;
    }
}
