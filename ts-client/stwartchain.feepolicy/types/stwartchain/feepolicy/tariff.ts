/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Fees } from "./fees";

export const protobufPackage = "stwartchain.feepolicy";

export interface Tariff {
  denom: string;
  id: number;
  amount: string;
  minRefBalance: string;
  fees: Fees[];
}

function createBaseTariff(): Tariff {
  return { denom: "", id: 0, amount: "", minRefBalance: "", fees: [] };
}

export const Tariff = {
  encode(message: Tariff, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    if (message.minRefBalance !== "") {
      writer.uint32(34).string(message.minRefBalance);
    }
    for (const v of message.fees) {
      Fees.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Tariff {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTariff();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.denom = reader.string();
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

          message.amount = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.minRefBalance = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.fees.push(Fees.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Tariff {
    return {
      denom: isSet(object.denom) ? String(object.denom) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      amount: isSet(object.amount) ? String(object.amount) : "",
      minRefBalance: isSet(object.minRefBalance) ? String(object.minRefBalance) : "",
      fees: Array.isArray(object?.fees) ? object.fees.map((e: any) => Fees.fromJSON(e)) : [],
    };
  },

  toJSON(message: Tariff): unknown {
    const obj: any = {};
    if (message.denom !== "") {
      obj.denom = message.denom;
    }
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    if (message.amount !== "") {
      obj.amount = message.amount;
    }
    if (message.minRefBalance !== "") {
      obj.minRefBalance = message.minRefBalance;
    }
    if (message.fees?.length) {
      obj.fees = message.fees.map((e) => Fees.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Tariff>, I>>(base?: I): Tariff {
    return Tariff.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Tariff>, I>>(object: I): Tariff {
    const message = createBaseTariff();
    message.denom = object.denom ?? "";
    message.id = object.id ?? 0;
    message.amount = object.amount ?? "";
    message.minRefBalance = object.minRefBalance ?? "";
    message.fees = object.fees?.map((e) => Fees.fromPartial(e)) || [];
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
