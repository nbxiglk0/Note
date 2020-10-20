unit Unit1;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ExtCtrls,
  ComCtrls;

type

  { TPeach }

  TPeach = class(TForm)
    Home: TButton;
    PageControl1: TPageControl;
    Scan: TButton;
    Button3: TButton;
    procedure FormCreate(Sender: TObject);
    procedure HomeClick(Sender: TObject);
  private

  public

  end;

var
  Peach: TPeach;

implementation

{$R *.lfm}

{ TPeach }

procedure TPeach.FormCreate(Sender: TObject);
begin

end;

procedure TPeach.HomeClick(Sender: TObject);
begin

end;

end.

