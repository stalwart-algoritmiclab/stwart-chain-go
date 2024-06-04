/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "stwartchain.stats";

export interface AssetDailyStats {
  /** fee policy module */
  amountWithFee: Coin[];
  amountNoFee: Coin[];
  fee: Coin[];
  countWithFee: number;
  countNoFee: number;
  /** core module */
  burned: Coin[];
  countBurned: number;
  issued: Coin[];
  countIssued: number;
  withdraw: Coin[];
  countWithdraw: number;
  refReward: Coin[];
  countRefReward: number;
  sysRefReward: Coin[];
  countSysRefReward: number;
}

function createBaseAssetDailyStats(): AssetDailyStats {
  return {
    amountWithFee: [],
    amountNoFee: [],
    fee: [],
    countWithFee: 0,
    countNoFee: 0,
    burned: [],
    countBurned: 0,
    issued: [],
    countIssued: 0,
    withdraw: [],
    countWithdraw: 0,
    refReward: [],
    countRefReward: 0,
    sysRefReward: [],
    countSysRefReward: 0,
  };
}

export const AssetDailyStats = {
  encode(message: AssetDailyStats, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.amountWithFee) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.amountNoFee) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.fee) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.countWithFee !== 0) {
      writer.uint32(32).int32(message.countWithFee);
    }
    if (message.countNoFee !== 0) {
      writer.uint32(40).int32(message.countNoFee);
    }
    for (const v of message.burned) {
      Coin.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    if (message.countBurned !== 0) {
      writer.uint32(56).uint64(message.countBurned);
    }
    for (const v of message.issued) {
      Coin.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.countIssued !== 0) {
      writer.uint32(72).uint64(message.countIssued);
    }
    for (const v of message.withdraw) {
      Coin.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    if (message.countWithdraw !== 0) {
      writer.uint32(88).uint64(message.countWithdraw);
    }
    for (const v of message.refReward) {
      Coin.encode(v!, writer.uint32(98).fork()).ldelim();
    }
    if (message.countRefReward !== 0) {
      writer.uint32(104).uint64(message.countRefReward);
    }
    for (const v of message.sysRefReward) {
      Coin.encode(v!, writer.uint32(130).fork()).ldelim();
    }
    if (message.countSysRefReward !== 0) {
      writer.uint32(136).uint64(message.countSysRefReward);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AssetDailyStats {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAssetDailyStats();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.amountWithFee.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.amountNoFee.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.fee.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.countWithFee = reader.int32();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.countNoFee = reader.int32();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.burned.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.countBurned = longToNumber(reader.uint64() as Long);
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.issued.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 9:
          if (tag !== 72) {
            break;
          }

          message.countIssued = longToNumber(reader.uint64() as Long);
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.withdraw.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 11:
          if (tag !== 88) {
            break;
          }

          message.countWithdraw = longToNumber(reader.uint64() as Long);
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.refReward.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 13:
          if (tag !== 104) {
            break;
          }

          message.countRefReward = longToNumber(reader.uint64() as Long);
          continue;
        case 16:
          if (tag !== 130) {
            break;
          }

          message.sysRefReward.push(Coin.decode(reader, reader.uint32()));
          continue;
        case 17:
          if (tag !== 136) {
            break;
          }

          message.countSysRefReward = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AssetDailyStats {
    return {
      amountWithFee: Array.isArray(object?.amountWithFee) ? object.amountWithFee.map((e: any) => Coin.fromJSON(e)) : [],
      amountNoFee: Array.isArray(object?.amountNoFee) ? object.amountNoFee.map((e: any) => Coin.fromJSON(e)) : [],
      fee: Array.isArray(object?.fee) ? object.fee.map((e: any) => Coin.fromJSON(e)) : [],
      countWithFee: isSet(object.countWithFee) ? Number(object.countWithFee) : 0,
      countNoFee: isSet(object.countNoFee) ? Number(object.countNoFee) : 0,
      burned: Array.isArray(object?.burned) ? object.burned.map((e: any) => Coin.fromJSON(e)) : [],
      countBurned: isSet(object.countBurned) ? Number(object.countBurned) : 0,
      issued: Array.isArray(object?.issued) ? object.issued.map((e: any) => Coin.fromJSON(e)) : [],
      countIssued: isSet(object.countIssued) ? Number(object.countIssued) : 0,
      withdraw: Array.isArray(object?.withdraw) ? object.withdraw.map((e: any) => Coin.fromJSON(e)) : [],
      countWithdraw: isSet(object.countWithdraw) ? Number(object.countWithdraw) : 0,
      refReward: Array.isArray(object?.refReward) ? object.refReward.map((e: any) => Coin.fromJSON(e)) : [],
      countRefReward: isSet(object.countRefReward) ? Number(object.countRefReward) : 0,
      sysRefReward: Array.isArray(object?.sysRefReward) ? object.sysRefReward.map((e: any) => Coin.fromJSON(e)) : [],
      countSysRefReward: isSet(object.countSysRefReward) ? Number(object.countSysRefReward) : 0,
    };
  },

  toJSON(message: AssetDailyStats): unknown {
    const obj: any = {};
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
    if (message.burned?.length) {
      obj.burned = message.burned.map((e) => Coin.toJSON(e));
    }
    if (message.countBurned !== 0) {
      obj.countBurned = Math.round(message.countBurned);
    }
    if (message.issued?.length) {
      obj.issued = message.issued.map((e) => Coin.toJSON(e));
    }
    if (message.countIssued !== 0) {
      obj.countIssued = Math.round(message.countIssued);
    }
    if (message.withdraw?.length) {
      obj.withdraw = message.withdraw.map((e) => Coin.toJSON(e));
    }
    if (message.countWithdraw !== 0) {
      obj.countWithdraw = Math.round(message.countWithdraw);
    }
    if (message.refReward?.length) {
      obj.refReward = message.refReward.map((e) => Coin.toJSON(e));
    }
    if (message.countRefReward !== 0) {
      obj.countRefReward = Math.round(message.countRefReward);
    }
    if (message.sysRefReward?.length) {
      obj.sysRefReward = message.sysRefReward.map((e) => Coin.toJSON(e));
    }
    if (message.countSysRefReward !== 0) {
      obj.countSysRefReward = Math.round(message.countSysRefReward);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AssetDailyStats>, I>>(base?: I): AssetDailyStats {
    return AssetDailyStats.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<AssetDailyStats>, I>>(object: I): AssetDailyStats {
    const message = createBaseAssetDailyStats();
    message.amountWithFee = object.amountWithFee?.map((e) => Coin.fromPartial(e)) || [];
    message.amountNoFee = object.amountNoFee?.map((e) => Coin.fromPartial(e)) || [];
    message.fee = object.fee?.map((e) => Coin.fromPartial(e)) || [];
    message.countWithFee = object.countWithFee ?? 0;
    message.countNoFee = object.countNoFee ?? 0;
    message.burned = object.burned?.map((e) => Coin.fromPartial(e)) || [];
    message.countBurned = object.countBurned ?? 0;
    message.issued = object.issued?.map((e) => Coin.fromPartial(e)) || [];
    message.countIssued = object.countIssued ?? 0;
    message.withdraw = object.withdraw?.map((e) => Coin.fromPartial(e)) || [];
    message.countWithdraw = object.countWithdraw ?? 0;
    message.refReward = object.refReward?.map((e) => Coin.fromPartial(e)) || [];
    message.countRefReward = object.countRefReward ?? 0;
    message.sysRefReward = object.sysRefReward?.map((e) => Coin.fromPartial(e)) || [];
    message.countSysRefReward = object.countSysRefReward ?? 0;
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
