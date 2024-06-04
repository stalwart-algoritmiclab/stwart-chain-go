/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Params } from "./params";
import { Tariff } from "./tariff";

export const protobufPackage = "stwartchain.feepolicy";

/** MsgUpdateParams is the Msg/UpdateParams request type. */
export interface MsgUpdateParams {
  /** authority is the address that controls the module (defaults to x/gov unless overwritten). */
  authority: string;
  /** NOTE: All parameters must be supplied. */
  params: Params | undefined;
}

/**
 * MsgUpdateParamsResponse defines the response structure for executing a
 * MsgUpdateParams message.
 */
export interface MsgUpdateParamsResponse {
}

export interface MsgCreateAddresses {
  creator: string;
  address: string;
}

export interface MsgCreateAddressesResponse {
  id: number;
}

export interface MsgUpdateAddresses {
  creator: string;
  id: number;
  address: string;
}

export interface MsgUpdateAddressesResponse {
}

export interface MsgDeleteAddresses {
  creator: string;
  id: number;
}

export interface MsgDeleteAddressesResponse {
}

export interface MsgCreateTariffs {
  creator: string;
  denom: string;
  tariffs: Tariff | undefined;
}

export interface MsgCreateTariffsResponse {
}

export interface MsgUpdateTariffs {
  creator: string;
  denom: string;
  tariffs: Tariff | undefined;
}

export interface MsgUpdateTariffsResponse {
}

export interface MsgDeleteTariffs {
  creator: string;
  denom: string;
  tariffID: string;
  feeID: string;
}

export interface MsgDeleteTariffsResponse {
}

function createBaseMsgUpdateParams(): MsgUpdateParams {
  return { authority: "", params: undefined };
}

export const MsgUpdateParams = {
  encode(message: MsgUpdateParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateParams {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.authority = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.params = Params.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateParams {
    return {
      authority: isSet(object.authority) ? String(object.authority) : "",
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
    };
  },

  toJSON(message: MsgUpdateParams): unknown {
    const obj: any = {};
    if (message.authority !== "") {
      obj.authority = message.authority;
    }
    if (message.params !== undefined) {
      obj.params = Params.toJSON(message.params);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateParams>, I>>(base?: I): MsgUpdateParams {
    return MsgUpdateParams.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateParams>, I>>(object: I): MsgUpdateParams {
    const message = createBaseMsgUpdateParams();
    message.authority = object.authority ?? "";
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseMsgUpdateParamsResponse(): MsgUpdateParamsResponse {
  return {};
}

export const MsgUpdateParamsResponse = {
  encode(_: MsgUpdateParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateParamsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateParamsResponse {
    return {};
  },

  toJSON(_: MsgUpdateParamsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateParamsResponse>, I>>(base?: I): MsgUpdateParamsResponse {
    return MsgUpdateParamsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateParamsResponse>, I>>(_: I): MsgUpdateParamsResponse {
    const message = createBaseMsgUpdateParamsResponse();
    return message;
  },
};

function createBaseMsgCreateAddresses(): MsgCreateAddresses {
  return { creator: "", address: "" };
}

export const MsgCreateAddresses = {
  encode(message: MsgCreateAddresses, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAddresses {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAddresses();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.address = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAddresses {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgCreateAddresses): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.address !== "") {
      obj.address = message.address;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgCreateAddresses>, I>>(base?: I): MsgCreateAddresses {
    return MsgCreateAddresses.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgCreateAddresses>, I>>(object: I): MsgCreateAddresses {
    const message = createBaseMsgCreateAddresses();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgCreateAddressesResponse(): MsgCreateAddressesResponse {
  return { id: 0 };
}

export const MsgCreateAddressesResponse = {
  encode(message: MsgCreateAddressesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAddressesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAddressesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAddressesResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgCreateAddressesResponse): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgCreateAddressesResponse>, I>>(base?: I): MsgCreateAddressesResponse {
    return MsgCreateAddressesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgCreateAddressesResponse>, I>>(object: I): MsgCreateAddressesResponse {
    const message = createBaseMsgCreateAddressesResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgUpdateAddresses(): MsgUpdateAddresses {
  return { creator: "", id: 0, address: "" };
}

export const MsgUpdateAddresses = {
  encode(message: MsgUpdateAddresses, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateAddresses {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateAddresses();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.address = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateAddresses {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgUpdateAddresses): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    if (message.address !== "") {
      obj.address = message.address;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateAddresses>, I>>(base?: I): MsgUpdateAddresses {
    return MsgUpdateAddresses.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateAddresses>, I>>(object: I): MsgUpdateAddresses {
    const message = createBaseMsgUpdateAddresses();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgUpdateAddressesResponse(): MsgUpdateAddressesResponse {
  return {};
}

export const MsgUpdateAddressesResponse = {
  encode(_: MsgUpdateAddressesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateAddressesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateAddressesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateAddressesResponse {
    return {};
  },

  toJSON(_: MsgUpdateAddressesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateAddressesResponse>, I>>(base?: I): MsgUpdateAddressesResponse {
    return MsgUpdateAddressesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateAddressesResponse>, I>>(_: I): MsgUpdateAddressesResponse {
    const message = createBaseMsgUpdateAddressesResponse();
    return message;
  },
};

function createBaseMsgDeleteAddresses(): MsgDeleteAddresses {
  return { creator: "", id: 0 };
}

export const MsgDeleteAddresses = {
  encode(message: MsgDeleteAddresses, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAddresses {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAddresses();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteAddresses {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgDeleteAddresses): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgDeleteAddresses>, I>>(base?: I): MsgDeleteAddresses {
    return MsgDeleteAddresses.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgDeleteAddresses>, I>>(object: I): MsgDeleteAddresses {
    const message = createBaseMsgDeleteAddresses();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgDeleteAddressesResponse(): MsgDeleteAddressesResponse {
  return {};
}

export const MsgDeleteAddressesResponse = {
  encode(_: MsgDeleteAddressesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAddressesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAddressesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteAddressesResponse {
    return {};
  },

  toJSON(_: MsgDeleteAddressesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgDeleteAddressesResponse>, I>>(base?: I): MsgDeleteAddressesResponse {
    return MsgDeleteAddressesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgDeleteAddressesResponse>, I>>(_: I): MsgDeleteAddressesResponse {
    const message = createBaseMsgDeleteAddressesResponse();
    return message;
  },
};

function createBaseMsgCreateTariffs(): MsgCreateTariffs {
  return { creator: "", denom: "", tariffs: undefined };
}

export const MsgCreateTariffs = {
  encode(message: MsgCreateTariffs, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denom !== "") {
      writer.uint32(18).string(message.denom);
    }
    if (message.tariffs !== undefined) {
      Tariff.encode(message.tariffs, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateTariffs {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateTariffs();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.denom = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.tariffs = Tariff.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTariffs {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      denom: isSet(object.denom) ? String(object.denom) : "",
      tariffs: isSet(object.tariffs) ? Tariff.fromJSON(object.tariffs) : undefined,
    };
  },

  toJSON(message: MsgCreateTariffs): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.denom !== "") {
      obj.denom = message.denom;
    }
    if (message.tariffs !== undefined) {
      obj.tariffs = Tariff.toJSON(message.tariffs);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgCreateTariffs>, I>>(base?: I): MsgCreateTariffs {
    return MsgCreateTariffs.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgCreateTariffs>, I>>(object: I): MsgCreateTariffs {
    const message = createBaseMsgCreateTariffs();
    message.creator = object.creator ?? "";
    message.denom = object.denom ?? "";
    message.tariffs = (object.tariffs !== undefined && object.tariffs !== null)
      ? Tariff.fromPartial(object.tariffs)
      : undefined;
    return message;
  },
};

function createBaseMsgCreateTariffsResponse(): MsgCreateTariffsResponse {
  return {};
}

export const MsgCreateTariffsResponse = {
  encode(_: MsgCreateTariffsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateTariffsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateTariffsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): MsgCreateTariffsResponse {
    return {};
  },

  toJSON(_: MsgCreateTariffsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgCreateTariffsResponse>, I>>(base?: I): MsgCreateTariffsResponse {
    return MsgCreateTariffsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgCreateTariffsResponse>, I>>(_: I): MsgCreateTariffsResponse {
    const message = createBaseMsgCreateTariffsResponse();
    return message;
  },
};

function createBaseMsgUpdateTariffs(): MsgUpdateTariffs {
  return { creator: "", denom: "", tariffs: undefined };
}

export const MsgUpdateTariffs = {
  encode(message: MsgUpdateTariffs, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denom !== "") {
      writer.uint32(18).string(message.denom);
    }
    if (message.tariffs !== undefined) {
      Tariff.encode(message.tariffs, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateTariffs {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateTariffs();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.denom = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.tariffs = Tariff.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateTariffs {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      denom: isSet(object.denom) ? String(object.denom) : "",
      tariffs: isSet(object.tariffs) ? Tariff.fromJSON(object.tariffs) : undefined,
    };
  },

  toJSON(message: MsgUpdateTariffs): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.denom !== "") {
      obj.denom = message.denom;
    }
    if (message.tariffs !== undefined) {
      obj.tariffs = Tariff.toJSON(message.tariffs);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateTariffs>, I>>(base?: I): MsgUpdateTariffs {
    return MsgUpdateTariffs.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateTariffs>, I>>(object: I): MsgUpdateTariffs {
    const message = createBaseMsgUpdateTariffs();
    message.creator = object.creator ?? "";
    message.denom = object.denom ?? "";
    message.tariffs = (object.tariffs !== undefined && object.tariffs !== null)
      ? Tariff.fromPartial(object.tariffs)
      : undefined;
    return message;
  },
};

function createBaseMsgUpdateTariffsResponse(): MsgUpdateTariffsResponse {
  return {};
}

export const MsgUpdateTariffsResponse = {
  encode(_: MsgUpdateTariffsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateTariffsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateTariffsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateTariffsResponse {
    return {};
  },

  toJSON(_: MsgUpdateTariffsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateTariffsResponse>, I>>(base?: I): MsgUpdateTariffsResponse {
    return MsgUpdateTariffsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateTariffsResponse>, I>>(_: I): MsgUpdateTariffsResponse {
    const message = createBaseMsgUpdateTariffsResponse();
    return message;
  },
};

function createBaseMsgDeleteTariffs(): MsgDeleteTariffs {
  return { creator: "", denom: "", tariffID: "", feeID: "" };
}

export const MsgDeleteTariffs = {
  encode(message: MsgDeleteTariffs, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denom !== "") {
      writer.uint32(18).string(message.denom);
    }
    if (message.tariffID !== "") {
      writer.uint32(26).string(message.tariffID);
    }
    if (message.feeID !== "") {
      writer.uint32(34).string(message.feeID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteTariffs {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteTariffs();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.denom = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.tariffID = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.feeID = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteTariffs {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      denom: isSet(object.denom) ? String(object.denom) : "",
      tariffID: isSet(object.tariffID) ? String(object.tariffID) : "",
      feeID: isSet(object.feeID) ? String(object.feeID) : "",
    };
  },

  toJSON(message: MsgDeleteTariffs): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.denom !== "") {
      obj.denom = message.denom;
    }
    if (message.tariffID !== "") {
      obj.tariffID = message.tariffID;
    }
    if (message.feeID !== "") {
      obj.feeID = message.feeID;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgDeleteTariffs>, I>>(base?: I): MsgDeleteTariffs {
    return MsgDeleteTariffs.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgDeleteTariffs>, I>>(object: I): MsgDeleteTariffs {
    const message = createBaseMsgDeleteTariffs();
    message.creator = object.creator ?? "";
    message.denom = object.denom ?? "";
    message.tariffID = object.tariffID ?? "";
    message.feeID = object.feeID ?? "";
    return message;
  },
};

function createBaseMsgDeleteTariffsResponse(): MsgDeleteTariffsResponse {
  return {};
}

export const MsgDeleteTariffsResponse = {
  encode(_: MsgDeleteTariffsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteTariffsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteTariffsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteTariffsResponse {
    return {};
  },

  toJSON(_: MsgDeleteTariffsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgDeleteTariffsResponse>, I>>(base?: I): MsgDeleteTariffsResponse {
    return MsgDeleteTariffsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgDeleteTariffsResponse>, I>>(_: I): MsgDeleteTariffsResponse {
    const message = createBaseMsgDeleteTariffsResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /**
   * UpdateParams defines a (governance) operation for updating the module
   * parameters. The authority defaults to the x/gov module account.
   */
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse>;
  CreateAddresses(request: MsgCreateAddresses): Promise<MsgCreateAddressesResponse>;
  UpdateAddresses(request: MsgUpdateAddresses): Promise<MsgUpdateAddressesResponse>;
  DeleteAddresses(request: MsgDeleteAddresses): Promise<MsgDeleteAddressesResponse>;
  CreateTariffs(request: MsgCreateTariffs): Promise<MsgCreateTariffsResponse>;
  UpdateTariffs(request: MsgUpdateTariffs): Promise<MsgUpdateTariffsResponse>;
  DeleteTariffs(request: MsgDeleteTariffs): Promise<MsgDeleteTariffsResponse>;
}

export const MsgServiceName = "stwartchain.feepolicy.Msg";
export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || MsgServiceName;
    this.rpc = rpc;
    this.UpdateParams = this.UpdateParams.bind(this);
    this.CreateAddresses = this.CreateAddresses.bind(this);
    this.UpdateAddresses = this.UpdateAddresses.bind(this);
    this.DeleteAddresses = this.DeleteAddresses.bind(this);
    this.CreateTariffs = this.CreateTariffs.bind(this);
    this.UpdateTariffs = this.UpdateTariffs.bind(this);
    this.DeleteTariffs = this.DeleteTariffs.bind(this);
  }
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse> {
    const data = MsgUpdateParams.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateParams", data);
    return promise.then((data) => MsgUpdateParamsResponse.decode(_m0.Reader.create(data)));
  }

  CreateAddresses(request: MsgCreateAddresses): Promise<MsgCreateAddressesResponse> {
    const data = MsgCreateAddresses.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateAddresses", data);
    return promise.then((data) => MsgCreateAddressesResponse.decode(_m0.Reader.create(data)));
  }

  UpdateAddresses(request: MsgUpdateAddresses): Promise<MsgUpdateAddressesResponse> {
    const data = MsgUpdateAddresses.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateAddresses", data);
    return promise.then((data) => MsgUpdateAddressesResponse.decode(_m0.Reader.create(data)));
  }

  DeleteAddresses(request: MsgDeleteAddresses): Promise<MsgDeleteAddressesResponse> {
    const data = MsgDeleteAddresses.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteAddresses", data);
    return promise.then((data) => MsgDeleteAddressesResponse.decode(_m0.Reader.create(data)));
  }

  CreateTariffs(request: MsgCreateTariffs): Promise<MsgCreateTariffsResponse> {
    const data = MsgCreateTariffs.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateTariffs", data);
    return promise.then((data) => MsgCreateTariffsResponse.decode(_m0.Reader.create(data)));
  }

  UpdateTariffs(request: MsgUpdateTariffs): Promise<MsgUpdateTariffsResponse> {
    const data = MsgUpdateTariffs.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateTariffs", data);
    return promise.then((data) => MsgUpdateTariffsResponse.decode(_m0.Reader.create(data)));
  }

  DeleteTariffs(request: MsgDeleteTariffs): Promise<MsgDeleteTariffsResponse> {
    const data = MsgDeleteTariffs.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteTariffs", data);
    return promise.then((data) => MsgDeleteTariffsResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
