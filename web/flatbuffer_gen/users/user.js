"use strict";
// automatically generated by the FlatBuffers compiler, do not modify
Object.defineProperty(exports, "__esModule", { value: true });
exports.User = void 0;
var flatbuffers = require("flatbuffers");
var address_js_1 = require("../users/address.js");
var User = /** @class */ (function () {
    function User() {
        this.bb = null;
        this.bb_pos = 0;
    }
    User.prototype.__init = function (i, bb) {
        this.bb_pos = i;
        this.bb = bb;
        return this;
    };
    User.getRootAsUser = function (bb, obj) {
        return (obj || new User()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
    };
    User.getSizePrefixedRootAsUser = function (bb, obj) {
        bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
        return (obj || new User()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
    };
    User.prototype.name = function (optionalEncoding) {
        var offset = this.bb.__offset(this.bb_pos, 4);
        return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
    };
    User.prototype.email = function (optionalEncoding) {
        var offset = this.bb.__offset(this.bb_pos, 6);
        return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
    };
    User.prototype.address = function (index, obj) {
        var offset = this.bb.__offset(this.bb_pos, 8);
        return offset ? (obj || new address_js_1.Address()).__init(this.bb.__indirect(this.bb.__vector(this.bb_pos + offset) + index * 4), this.bb) : null;
    };
    User.prototype.addressLength = function () {
        var offset = this.bb.__offset(this.bb_pos, 8);
        return offset ? this.bb.__vector_len(this.bb_pos + offset) : 0;
    };
    User.prototype.phone = function (optionalEncoding) {
        var offset = this.bb.__offset(this.bb_pos, 10);
        return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
    };
    User.startUser = function (builder) {
        builder.startObject(4);
    };
    User.addName = function (builder, nameOffset) {
        builder.addFieldOffset(0, nameOffset, 0);
    };
    User.addEmail = function (builder, emailOffset) {
        builder.addFieldOffset(1, emailOffset, 0);
    };
    User.addAddress = function (builder, addressOffset) {
        builder.addFieldOffset(2, addressOffset, 0);
    };
    User.createAddressVector = function (builder, data) {
        builder.startVector(4, data.length, 4);
        for (var i = data.length - 1; i >= 0; i--) {
            builder.addOffset(data[i]);
        }
        return builder.endVector();
    };
    User.startAddressVector = function (builder, numElems) {
        builder.startVector(4, numElems, 4);
    };
    User.addPhone = function (builder, phoneOffset) {
        builder.addFieldOffset(3, phoneOffset, 0);
    };
    User.endUser = function (builder) {
        var offset = builder.endObject();
        return offset;
    };
    User.finishUserBuffer = function (builder, offset) {
        builder.finish(offset);
    };
    User.finishSizePrefixedUserBuffer = function (builder, offset) {
        builder.finish(offset, undefined, true);
    };
    User.createUser = function (builder, nameOffset, emailOffset, addressOffset, phoneOffset) {
        User.startUser(builder);
        User.addName(builder, nameOffset);
        User.addEmail(builder, emailOffset);
        User.addAddress(builder, addressOffset);
        User.addPhone(builder, phoneOffset);
        return User.endUser(builder);
    };
    return User;
}());
exports.User = User;