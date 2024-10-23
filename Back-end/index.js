"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var express_1 = require("express");
var api = (0, express_1.default)();
api.listen(3000, function () {
    console.log('Servidor Conectado! \nHost http://localhost:3000/');
});
