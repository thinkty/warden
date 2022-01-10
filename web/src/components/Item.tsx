import React from 'react';

type RecordContent = {
  String: string;
  Valid: boolean;
}

export type Record = {
  Id: number;
  Date: string; // TODO: Date ?
  Beacon: string;
  Name: string;
  RecordType: number;
  Record: RecordContent;
};

type Props = {
  record: Record;
} & typeof defaultProps;

const defaultProps = {};

export const Item = ({ record }: Props): JSX.Element => {

  return (
    <div
      onClick={() => { console.log(record) }}
      style={{
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
        alignItems: 'center',
        border: 'thin solid white',
        overflow: 'auto',
        cursor: 'pointer',
      }}
    >
      {/* TODO: Handle overflowing */}
      {/* TODO: Update props and display accordingly */}
      {/* Item Header */}
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <div
          style={{
            fontSize: 20,
          }}
        >
          { record.Name }
        </div>
        <div
          style={{
            fontSize: 13,
          }}
        >
          { record.Beacon }
        </div>
      </div>
      {/* Item Content */}
      <div
        style={{
          fontSize: 44,
          fontWeight: 'bold',
        }}
      >
        { record.Record.String }
      </div>
      {/* Item Footer */}
      <div
        style={{

        }}
      />
    </div>
  );
}
Item.defaultProps = defaultProps;
