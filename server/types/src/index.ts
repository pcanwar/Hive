import express, { Express, Request, Response } from "express";
import body_parser from "body-parser"
import mongoose, {ConnectOptions} from "mongoose";



const app: Express = express();
const port :Number = 4000

const username: String = "admain"
const password: String = "123"
const host: string = "mongo"
const database:string = "newsDB"

const mongoURI = ``

mongoose.connect(mongoURI, {

}).then(() => {
    console.log("Connected");
}).catch((err) => {
    console.error("Error", err);
})

app.use(express.json())
// app.use("./client", client)

app.get('/', (req: Request, res: Response) => {
    res.send({
        "msg": "Conne"
    })
})

app.use(body_parser.urlencoded({
    extended: false,
    limit: "50mb"
}))
app.use(body_parser.json({
    limit: "50mb"
}))

app.listen(port, () => {
    console.log("listening on port " + port)
})