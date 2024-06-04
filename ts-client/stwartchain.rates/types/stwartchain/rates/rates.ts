/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stwartchain.rates";

export interface Rates {
  denom: string;
  rate: number;
  creator: string;
  decimals: number;
}

function createBaseRates(): Rates {
  return { denom: "", rate: 0, creator: "", decimals: 0 };
}

export const Rates = {
  encode(message: Rates, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    if (message.rate !== 0) {
      writer.uint32(17).double(message.rate);
    }
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
    }
    if (message.decimals !== 0) {
      writer.uint32(32).int32(message.decimals);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Rates {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRates();
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
          if (tag !== 17) {
            break;
          }

          message.rate = reader.double();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.decimals = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Rates {
    return {
      denom: isSet(object.denom) ? String(object.denom) : "",
      rate: isSet(object.rate) ? Number(object.rate) : 0,
      creator: isSet(object.creator) ? String(object.creator) : "",
      decimals: isSet(object.decimals) ? Number(object.decimals) : 0,
    };
  },

  toJSON(message: Rates): unknown {
    const obj: any = {};
    if (message.denom !== "") {
      obj.denom = message.denom;
    }
    if (message.rate !== 0) {
      obj.rate = message.rate;
    }
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.decimals !== 0) {
      obj.decimals = Math.round(message.decimals);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Rates>, I>>(base?: I): Rates {
    return Rates.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Rates>, I>>(object: I): Rates {
    const message = createBaseRates();
    message.denom = object.denom ?? "";
    message.rate = object.rate ?? 0;
    message.creator = object.creator ?? "";
    message.decimals = object.decimals ?? 0;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
