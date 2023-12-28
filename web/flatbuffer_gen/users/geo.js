"use strict";
// automatically generated by the FlatBuffers compiler, do not modify
Object.defineProperty(exports, "__esModule", { value: true });
exports.Geo = void 0;
var flatbuffers = require("flatbuffers");
var Geo = /** @class */ (function () {
    function Geo() {
        this.bb = null;
        this.bb_pos = 0;
    }
    Geo.prototype.__init = function (i, bb) {
        this.bb_pos = i;
        this.bb = bb;
        return this;
    };
    Geo.getRootAsGeo = function (bb, obj) {
        return (obj || new Geo()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
    };
    Geo.getSizePrefixedRootAsGeo = function (bb, obj) {
        bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
        return (obj || new Geo()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
    };
    Geo.prototype.lat = function (optionalEncoding) {
        var offset = this.bb.__offset(this.bb_pos, 4);
        return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
    };
    Geo.prototype.long = function (optionalEncoding) {
        var offset = this.bb.__offset(this.bb_pos, 6);
        return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
    };
    Geo.startGeo = function (builder) {
        builder.startObject(2);
    };
    Geo.addLat = function (builder, latOffset) {
        builder.addFieldOffset(0, latOffset, 0);
    };
    Geo.addLong = function (builder, longOffset) {
        builder.addFieldOffset(1, longOffset, 0);
    };
    Geo.endGeo = function (builder) {
        var offset = builder.endObject();
        return offset;
    };
    Geo.createGeo = function (builder, latOffset, longOffset) {
        Geo.startGeo(builder);
        Geo.addLat(builder, latOffset);
        Geo.addLong(builder, longOffset);
        return Geo.endGeo(builder);
    };
    return Geo;
}());
exports.Geo = Geo;
