using aspnetcoreapp;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddControllers();
builder.Services.AddGrpc(options =>
{
    options.EnableDetailedErrors = true;
});
var app = builder.Build();

app.MapGrpcService<GreeterService>();
app.MapControllers();
app.Run();
