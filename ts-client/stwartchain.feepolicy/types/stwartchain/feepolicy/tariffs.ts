/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Tariff } from "./tariff";

export const protobufPackage = "stwartchain.feepolicy";

export interface Tariffs {
  denom: string;
  tariffs: Tariff[];
  creator: string;
}

function createBaseTariffs(): Tariffs {
  return { denom: "", tariffs: [], creator: "" };
}

export const Tariffs = {
  encode(message: Tariffs, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    for (const v of message.tariffs) {
      Tariff.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Tariffs {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTariffs();
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
          if (tag !== 18) {
            break;
          }

          message.tariffs.push(Tariff.decode(reader, reader.uint32()));
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.creator = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Tariffs {
    return {
      denom: isSet(object.denom) ? String(object.denom) : "",
      tariffs: Array.isArray(object?.tariffs) ? object.tariffs.map((e: any) => Tariff.fromJSON(e)) : [],
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: Tariffs): unknown {
    const obj: any = {};
    if (message.denom !== "") {
      obj.denom = message.denom;
    }
    if (message.tariffs?.length) {
      obj.tariffs = message.tariffs.map((e) => Tariff.toJSON(e));
    }
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Tariffs>, I>>(base?: I): Tariffs {
    return Tariffs.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Tariffs>, I>>(object: I): Tariffs {
    const message = createBaseTariffs();
    message.denom = object.denom ?? "";
    message.tariffs = object.tariffs?.map((e) => Tariff.fromPartial(e)) || [];
    message.creator = object.creator ?? "";
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
