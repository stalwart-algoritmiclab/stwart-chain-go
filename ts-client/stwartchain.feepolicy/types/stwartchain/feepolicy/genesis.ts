/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Address } from "./addresses";
import { Params } from "./params";
import { Tariff } from "./tariff";
import { Tariffs } from "./tariffs";

export const protobufPackage = "stwartchain.feepolicy";

/** GenesisState defines the feepolicy module's genesis state. */
export interface GenesisState {
  /** params defines all the parameters of the module. */
  params: Params | undefined;
  addressesList: Address[];
  addressesCount: number;
  tariffList: Tariff[];
  tariffsList: Tariffs[];
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, addressesList: [], addressesCount: 0, tariffList: [], tariffsList: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.addressesList) {
      Address.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.addressesCount !== 0) {
      writer.uint32(24).uint64(message.addressesCount);
    }
    for (const v of message.tariffList) {
      Tariff.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.tariffsList) {
      Tariffs.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.params = Params.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.addressesList.push(Address.decode(reader, reader.uint32()));
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.addressesCount = longToNumber(reader.uint64() as Long);
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.tariffList.push(Tariff.decode(reader, reader.uint32()));
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.tariffsList.push(Tariffs.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      addressesList: Array.isArray(object?.addressesList)
        ? object.addressesList.map((e: any) => Address.fromJSON(e))
        : [],
      addressesCount: isSet(object.addressesCount) ? Number(object.addressesCount) : 0,
      tariffList: Array.isArray(object?.tariffList) ? object.tariffList.map((e: any) => Tariff.fromJSON(e)) : [],
      tariffsList: Array.isArray(object?.tariffsList) ? object.tariffsList.map((e: any) => Tariffs.fromJSON(e)) : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.params !== undefined) {
      obj.params = Params.toJSON(message.params);
    }
    if (message.addressesList?.length) {
      obj.addressesList = message.addressesList.map((e) => Address.toJSON(e));
    }
    if (message.addressesCount !== 0) {
      obj.addressesCount = Math.round(message.addressesCount);
    }
    if (message.tariffList?.length) {
      obj.tariffList = message.tariffList.map((e) => Tariff.toJSON(e));
    }
    if (message.tariffsList?.length) {
      obj.tariffsList = message.tariffsList.map((e) => Tariffs.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GenesisState>, I>>(base?: I): GenesisState {
    return GenesisState.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.addressesList = object.addressesList?.map((e) => Address.fromPartial(e)) || [];
    message.addressesCount = object.addressesCount ?? 0;
    message.tariffList = object.tariffList?.map((e) => Tariff.fromPartial(e)) || [];
    message.tariffsList = object.tariffsList?.map((e) => Tariffs.fromPartial(e)) || [];
    return message;
  },
};

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
