/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stwartchain.stats";

export interface UserStats {
  countUniqueActiveUsers: number;
  countNewUsers: number;
}

function createBaseUserStats(): UserStats {
  return { countUniqueActiveUsers: 0, countNewUsers: 0 };
}

export const UserStats = {
  encode(message: UserStats, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.countUniqueActiveUsers !== 0) {
      writer.uint32(8).uint64(message.countUniqueActiveUsers);
    }
    if (message.countNewUsers !== 0) {
      writer.uint32(16).uint64(message.countNewUsers);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserStats {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserStats();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.countUniqueActiveUsers = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.countNewUsers = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UserStats {
    return {
      countUniqueActiveUsers: isSet(object.countUniqueActiveUsers) ? Number(object.countUniqueActiveUsers) : 0,
      countNewUsers: isSet(object.countNewUsers) ? Number(object.countNewUsers) : 0,
    };
  },

  toJSON(message: UserStats): unknown {
    const obj: any = {};
    if (message.countUniqueActiveUsers !== 0) {
      obj.countUniqueActiveUsers = Math.round(message.countUniqueActiveUsers);
    }
    if (message.countNewUsers !== 0) {
      obj.countNewUsers = Math.round(message.countNewUsers);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UserStats>, I>>(base?: I): UserStats {
    return UserStats.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UserStats>, I>>(object: I): UserStats {
    const message = createBaseUserStats();
    message.countUniqueActiveUsers = object.countUniqueActiveUsers ?? 0;
    message.countNewUsers = object.countNewUsers ?? 0;
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
