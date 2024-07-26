const mongoose = require('mongoose')
const DataSchema = new mongoose.Schema({
    data: {
        type: Object,
        required: false
    },
    createdAt: { type: Date, default: Date.now },
})
const StreetModel = mongoose.model('data', DataSchema)
module.exports = StreetModel
