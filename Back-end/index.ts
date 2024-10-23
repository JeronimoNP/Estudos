import Express  from "express";
import Cors from "cors";

const api = Express();
const cors = Cors();

api.use(bodyParser.json());


api.listen(3000, () => {
    console.log('Servidor Conectado! \nHost http://localhost:3000/');
});