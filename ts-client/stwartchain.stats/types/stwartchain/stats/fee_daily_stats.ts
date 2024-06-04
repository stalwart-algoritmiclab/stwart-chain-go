/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "stwartchain.stats";

export interface FeeDailyStats {
  id: number;
  amountWithFee: Coin[];
  amountNoFee: Coin[];
  fee: Coin[];
  countWithFee: number;
  countNoFee: number;
}

function createBaseFeeDailyStats(): FeeDailyStats {
  return { id: 0, amountWithFee: [], amountNoFee: [], fee: [], countWithFee: 0, countNoFee: 0 };
}

export const FeeDailyStats = {
  encode(message: FeeDailyStats, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    for (const v of message.amountWithFee) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.amountNoFee) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.fee) {
      Coin.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.countWithFee !== 0) {
      writer.uint32(40).int32(message.countWithFee);
    }
    if (message.countNoFee !== 0) {
      writer.uint32(48).int32(message.countNoFee);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeeDailyStats {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeeDailyStats();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.amountWithFee.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.amountNoFee.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.fee.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.countWithFee = reader.int32();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.countNoFee = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FeeDailyStats {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      amountWithFee: Array.isArray(object?.amountWithFee) ? object.amountWithFee.map((e: any) => Coin.fromJSON(e)) : [],
      amountNoFee: Array.isArray(object?.amountNoFee) ? object.amountNoFee.map((e: any) => Coin.fromJSON(e)) : [],
      fee: Array.isArray(object?.fee) ? object.fee.map((e: any) => Coin.fromJSON(e)) : [],
      countWithFee: isSet(object.countWithFee) ? Number(object.countWithFee) : 0,
      countNoFee: isSet(object.countNoFee) ? Number(object.countNoFee) : 0,
    };
  },

  toJSON(message: FeeDailyStats): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    if (message.amountWithFee?.length) {
      obj.amountWithFee = message.amountWithFee.map((e) => Coin.toJSON(e));
    }
    if (message.amountNoFee?.length) {
      obj.amountNoFee = message.amountNoFee.map((e) => Coin.toJSON(e));
    }
    if (message.fee?.length) {
      obj.fee = message.fee.map((e) => Coin.toJSON(e));
    }
    if (message.countWithFee !== 0) {
      obj.countWithFee = Math.round(message.countWithFee);
    }
    if (message.countNoFee !== 0) {
      obj.countNoFee = Math.round(message.countNoFee);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<FeeDailyStats>, I>>(base?: I): FeeDailyStats {
    return FeeDailyStats.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<FeeDailyStats>, I>>(object: I): FeeDailyStats {
    const message = createBaseFeeDailyStats();
    message.id = object.id ?? 0;
    message.amountWithFee = object.amountWithFee?.map((e) => Coin.fromPartial(e)) || [];
    message.amountNoFee = object.amountNoFee?.map((e) => Coin.fromPartial(e)) || [];
    message.fee = object.fee?.map((e) => Coin.fromPartial(e)) || [];
    message.countWithFee = object.countWithFee ?? 0;
    message.countNoFee = object.countNoFee ?? 0;
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
