module BitVector where
import Data.Text
import Data.Int
import Data.Bits










data Vect256 = Vect256 !Int64 !Int64 !Int64 !Int64 deriving (Eq)










xor256 :: Vect256 -> Vect256 -> Vect256
xor256 (Vect256 w0 x0 y0 z0) (Vect256 w1 x1 y1 z1) = Vect256 (xor w0 w1) (xor x0 x1) (xor y0 y1) (xor z0 z1)










and256 :: Vect256 -> Vect256 -> Vect256
and256 (Vect256 w0 x0 y0 z0) (Vect256 w1 x1 y1 z1) = Vect256 (w0 .&. w1) (x0 .&. x1) (y0 .&. y1) (z0 .&. z1)










or256 :: Vect256 -> Vect256 -> Vect256
or256 (Vect256 w0 x0 y0 z0) (Vect256 w1 x1 y1 z1) = Vect256 (w0 .|. w1) (x0 .|. x1) (y0 .|. y1) (z0 .|. z1)










not256 :: Vect256 -> Vect256
not256 (Vect256 w x y z) = Vect256 (complement w) (complement x) (complement y) (complement z)










popCount256 :: Vect256 -> Int
popCount256 (Vect256 w x y z) = (popCount w) + (popCount x) + (popCount y) + (popCount z)










match256 :: Vect256 -> Vect256 -> Int
match256 a b = popCount256 $ and256 a b










data SemanticObj =
  SemanticObj {
    name :: !Vect256,
    pos  :: !Vect256 }
