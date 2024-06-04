/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stwartchain.feepolicy";

export interface Fees {
  amountFrom: string;
  fee: string;
  refReward: string;
  stakeReward: string;
  minAmount: number;
  noRefReward: boolean;
  creator: string;
  id: number;
}

function createBaseFees(): Fees {
  return {
    amountFrom: "",
    fee: "",
    refReward: "",
    stakeReward: "",
    minAmount: 0,
    noRefReward: false,
    creator: "",
    id: 0,
  };
}

export const Fees = {
  encode(message: Fees, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.amountFrom !== "") {
      writer.uint32(10).string(message.amountFrom);
    }
    if (message.fee !== "") {
      writer.uint32(18).string(message.fee);
    }
    if (message.refReward !== "") {
      writer.uint32(26).string(message.refReward);
    }
    if (message.stakeReward !== "") {
      writer.uint32(34).string(message.stakeReward);
    }
    if (message.minAmount !== 0) {
      writer.uint32(40).uint64(message.minAmount);
    }
    if (message.noRefReward === true) {
      writer.uint32(48).bool(message.noRefReward);
    }
    if (message.creator !== "") {
      writer.uint32(58).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(64).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Fees {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFees();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.amountFrom = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.fee = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.refReward = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.stakeReward = reader.string();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.minAmount = longToNumber(reader.uint64() as Long);
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.noRefReward = reader.bool();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 8:
          if (tag !== 64) {
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

  fromJSON(object: any): Fees {
    return {
      amountFrom: isSet(object.amountFrom) ? String(object.amountFrom) : "",
      fee: isSet(object.fee) ? String(object.fee) : "",
      refReward: isSet(object.refReward) ? String(object.refReward) : "",
      stakeReward: isSet(object.stakeReward) ? String(object.stakeReward) : "",
      minAmount: isSet(object.minAmount) ? Number(object.minAmount) : 0,
      noRefReward: isSet(object.noRefReward) ? Boolean(object.noRefReward) : false,
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: Fees): unknown {
    const obj: any = {};
    if (message.amountFrom !== "") {
      obj.amountFrom = message.amountFrom;
    }
    if (message.fee !== "") {
      obj.fee = message.fee;
    }
    if (message.refReward !== "") {
      obj.refReward = message.refReward;
    }
    if (message.stakeReward !== "") {
      obj.stakeReward = message.stakeReward;
    }
    if (message.minAmount !== 0) {
      obj.minAmount = Math.round(message.minAmount);
    }
    if (message.noRefReward === true) {
      obj.noRefReward = message.noRefReward;
    }
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Fees>, I>>(base?: I): Fees {
    return Fees.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Fees>, I>>(object: I): Fees {
    const message = createBaseFees();
    message.amountFrom = object.amountFrom ?? "";
    message.fee = object.fee ?? "";
    message.refReward = object.refReward ?? "";
    message.stakeReward = object.stakeReward ?? "";
    message.minAmount = object.minAmount ?? 0;
    message.noRefReward = object.noRefReward ?? false;
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
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
