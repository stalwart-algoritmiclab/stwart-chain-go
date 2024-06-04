/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { FeeDailyStats } from "./fee_daily_stats";

export const protobufPackage = "stwartchain.stats";

export interface FeeStats {
  date: string;
  index: string;
  stats: FeeDailyStats | undefined;
}

function createBaseFeeStats(): FeeStats {
  return { date: "", index: "", stats: undefined };
}

export const FeeStats = {
  encode(message: FeeStats, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.date !== "") {
      writer.uint32(10).string(message.date);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.stats !== undefined) {
      FeeDailyStats.encode(message.stats, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeeStats {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeeStats();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.date = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.index = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.stats = FeeDailyStats.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FeeStats {
    return {
      date: isSet(object.date) ? String(object.date) : "",
      index: isSet(object.index) ? String(object.index) : "",
      stats: isSet(object.stats) ? FeeDailyStats.fromJSON(object.stats) : undefined,
    };
  },

  toJSON(message: FeeStats): unknown {
    const obj: any = {};
    if (message.date !== "") {
      obj.date = message.date;
    }
    if (message.index !== "") {
      obj.index = message.index;
    }
    if (message.stats !== undefined) {
      obj.stats = FeeDailyStats.toJSON(message.stats);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<FeeStats>, I>>(base?: I): FeeStats {
    return FeeStats.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<FeeStats>, I>>(object: I): FeeStats {
    const message = createBaseFeeStats();
    message.date = object.date ?? "";
    message.index = object.index ?? "";
    message.stats = (object.stats !== undefined && object.stats !== null)
      ? FeeDailyStats.fromPartial(object.stats)
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
