import Interpreter
import Test.Tasty
import Test.Tasty.Hspec

spec_interpreter :: Spec
spec_interpreter = do
  let code =
        [ -- LOAD_VAL 1
          LoadVal 1,
          -- WRITE_VAR 'x'
          WriteVar 'x',
          -- LOAD_VAL 2
          LoadVal 2,
          -- WRITE_VAR 'y'
          WriteVar 'y',
          -- READ_VAR 'x'
          ReadVar 'x',
          -- LOAD_VAL 1
          LoadVal 1,
          -- ADD
          Add,
          -- READ_VAR 'y'
          ReadVar 'y',
          -- MULTIPLY
          Multiply,
          -- RETURN_VALUE
          ReturnValue
        ]

  it "runByteCode" $ runByteCode code `shouldBe` Right 4

main :: IO ()
main = do
  tests <- testSpec "Bytecode Interpreter" spec_interpreter
  defaultMain tests