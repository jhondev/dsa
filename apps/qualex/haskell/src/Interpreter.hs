module Interpreter where

import qualified Data.Sequence as SeqM

data ByteCode
  = LoadVal Int
  | WriteVar Char
  | ReadVar Char
  | Add
  | Multiply
  | ReturnValue

data Variable = Variable
  { variable :: Maybe Char,
    value :: Int
  }

interpret :: [ByteCode] -> [Variable] -> Either String Variable
interpret [] [] = Left "Empty code"
interpret [] ops = Right (head ops)
interpret code@(instr : tail) ops =
  case instr of
    LoadVal val -> interpret tail (Variable {variable = Nothing, value = val} : ops)
    WriteVar var -> do
      let op = head ops
      interpret tail (Variable {variable = Just var, value = value op} : ops)
    ReadVar var -> do
      let op = filter (\op -> variable op == Just var) ops
      interpret tail (Variable {variable = Just var, value = value (head op)} : ops)
    Add -> do
      let opVals = take 2 ops
      let val = sum (map value opVals)
      interpret tail (Variable {variable = Nothing, value = val} : ops)
    Multiply -> do
      let opVals = take 2 ops
      let val = product (map value opVals)
      interpret tail (Variable {variable = Nothing, value = val} : ops)
    ReturnValue -> interpret tail (Variable {variable = Nothing, value = value (head ops)} : ops)

runByteCode :: [ByteCode] -> Either String Int
runByteCode code = do
  let result = interpret code []
  case result of
    Left error -> Left error
    Right val -> Right (value val)