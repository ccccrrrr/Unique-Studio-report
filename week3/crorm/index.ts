let mysql = require('mysql')
let connection = mysql.createConnecion({
    host : 'localhost',
    user : 'root',
    password : ' ',
})

connection.connect(function(err) {
    if (err) {
        console.error('error connecting: ' + err.stack);
        return;
    }

    console.log('connected as id ' + connection.threadId);
});