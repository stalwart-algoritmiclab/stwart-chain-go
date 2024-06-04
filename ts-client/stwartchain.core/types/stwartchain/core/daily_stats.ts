/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "stwartchain.core";

export interface DailyStats {
  issuedCoins: Coin[];
  countIssued: number;
  burnedCoins: Coin[];
  countBurned: number;
  withdrawCoins: Coin[];
  countWithdraw: number;
}

function createBaseDailyStats(): DailyStats {
  return { issuedCoins: [], countIssued: 0, burnedCoins: [], countBurned: 0, withdrawCoins: [], countWithdraw: 0 };
}

export const DailyStats = {
  encode(message: DailyStats, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.issuedCoins) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.countIssued !== 0) {
      writer.uint32(16).uint64(message.countIssued);
    }
    for (const v of message.burnedCoins) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.countBurned !== 0) {
      writer.uint32(32).uint64(message.countBurned);
    }
    for (const v of message.withdrawCoins) {
      Coin.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.countWithdraw !== 0) {
      writer.uint32(48).uint64(message.countWithdraw);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DailyStats {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDailyStats();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.issuedCoins.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.countIssued = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.burnedCoins.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.countBurned = longToNumber(reader.uint64() as Long);
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.withdrawCoins.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.countWithdraw = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DailyStats {
    return {
      issuedCoins: Array.isArray(object?.issuedCoins) ? object.issuedCoins.map((e: any) => Coin.fromJSON(e)) : [],
      countIssued: isSet(object.countIssued) ? Number(object.countIssued) : 0,
      burnedCoins: Array.isArray(object?.burnedCoins) ? object.burnedCoins.map((e: any) => Coin.fromJSON(e)) : [],
      countBurned: isSet(object.countBurned) ? Number(object.countBurned) : 0,
      withdrawCoins: Array.isArray(object?.withdrawCoins) ? object.withdrawCoins.map((e: any) => Coin.fromJSON(e)) : [],
      countWithdraw: isSet(object.countWithdraw) ? Number(object.countWithdraw) : 0,
    };
  },

  toJSON(message: DailyStats): unknown {
    const obj: any = {};
    if (message.issuedCoins?.length) {
      obj.issuedCoins = message.issuedCoins.map((e) => Coin.toJSON(e));
    }
    if (message.countIssued !== 0) {
      obj.countIssued = Math.round(message.countIssued);
    }
    if (message.burnedCoins?.length) {
      obj.burnedCoins = message.burnedCoins.map((e) => Coin.toJSON(e));
    }
    if (message.countBurned !== 0) {
      obj.countBurned = Math.round(message.countBurned);
    }
    if (message.withdrawCoins?.length) {
      obj.withdrawCoins = message.withdrawCoins.map((e) => Coin.toJSON(e));
    }
    if (message.countWithdraw !== 0) {
      obj.countWithdraw = Math.round(message.countWithdraw);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DailyStats>, I>>(base?: I): DailyStats {
    return DailyStats.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DailyStats>, I>>(object: I): DailyStats {
    const message = createBaseDailyStats();
    message.issuedCoins = object.issuedCoins?.map((e) => Coin.fromPartial(e)) || [];
    message.countIssued = object.countIssued ?? 0;
    message.burnedCoins = object.burnedCoins?.map((e) => Coin.fromPartial(e)) || [];
    message.countBurned = object.countBurned ?? 0;
    message.withdrawCoins = object.withdrawCoins?.map((e) => Coin.fromPartial(e)) || [];
    message.countWithdraw = object.countWithdraw ?? 0;
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
