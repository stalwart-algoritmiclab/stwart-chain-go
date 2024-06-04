/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "stwartchain.stake";

export interface Stake {
  address: string;
  sellAmount: Coin | undefined;
}

function createBaseStake(): Stake {
  return { address: "", sellAmount: undefined };
}

export const Stake = {
  encode(message: Stake, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.sellAmount !== undefined) {
      Coin.encode(message.sellAmount, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Stake {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStake();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.address = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.sellAmount = Coin.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Stake {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      sellAmount: isSet(object.sellAmount) ? Coin.fromJSON(object.sellAmount) : undefined,
    };
  },

  toJSON(message: Stake): unknown {
    const obj: any = {};
    if (message.address !== "") {
      obj.address = message.address;
    }
    if (message.sellAmount !== undefined) {
      obj.sellAmount = Coin.toJSON(message.sellAmount);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Stake>, I>>(base?: I): Stake {
    return Stake.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Stake>, I>>(object: I): Stake {
    const message = createBaseStake();
    message.address = object.address ?? "";
    message.sellAmount = (object.sellAmount !== undefined && object.sellAmount !== null)
      ? Coin.fromPartial(object.sellAmount)
      : undefined;
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
