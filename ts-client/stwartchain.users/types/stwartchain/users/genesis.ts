/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Params } from "./params";
import { Stats } from "./stats";
import { UniqueUsers } from "./unique_users";

export const protobufPackage = "stwartchain.users";

/** GenesisState defines the users module's genesis state. */
export interface GenesisState {
  /** params defines all the parameters of the module. */
  params: Params | undefined;
  statsList: Stats[];
  uniqueUsersList: UniqueUsers[];
  totalUsers: number;
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, statsList: [], uniqueUsersList: [], totalUsers: 0 };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.statsList) {
      Stats.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.uniqueUsersList) {
      UniqueUsers.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.totalUsers !== 0) {
      writer.uint32(32).uint64(message.totalUsers);
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

          message.statsList.push(Stats.decode(reader, reader.uint32()));
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.uniqueUsersList.push(UniqueUsers.decode(reader, reader.uint32()));
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.totalUsers = longToNumber(reader.uint64() as Long);
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
      statsList: Array.isArray(object?.statsList) ? object.statsList.map((e: any) => Stats.fromJSON(e)) : [],
      uniqueUsersList: Array.isArray(object?.uniqueUsersList)
        ? object.uniqueUsersList.map((e: any) => UniqueUsers.fromJSON(e))
        : [],
      totalUsers: isSet(object.totalUsers) ? Number(object.totalUsers) : 0,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.params !== undefined) {
      obj.params = Params.toJSON(message.params);
    }
    if (message.statsList?.length) {
      obj.statsList = message.statsList.map((e) => Stats.toJSON(e));
    }
    if (message.uniqueUsersList?.length) {
      obj.uniqueUsersList = message.uniqueUsersList.map((e) => UniqueUsers.toJSON(e));
    }
    if (message.totalUsers !== 0) {
      obj.totalUsers = Math.round(message.totalUsers);
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
    message.statsList = object.statsList?.map((e) => Stats.fromPartial(e)) || [];
    message.uniqueUsersList = object.uniqueUsersList?.map((e) => UniqueUsers.fromPartial(e)) || [];
    message.totalUsers = object.totalUsers ?? 0;
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
