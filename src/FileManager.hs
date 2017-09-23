module FileManager where
import System.IO
import Text.Regex.Base
import qualified Text.Regex.PCRE.String as P
import qualified Data.Text              as T
import qualified Data.Text.Array        as A










eitherMonadMapA :: Monad m => (a -> c) -> m (Either a b) -> m (Either c b)
eitherMonadMapA f m = fmap (eithermap f) m
  where eithermap :: (a -> c) -> Either a b -> Either c b
        eithermap f (Left  a) = Left  $ f a
        eithermap f (Right b) = Right b










eitherMonadMapB :: Monad m =>  (b -> c) -> m (Either a b) -> m (Either a c)
eitherMonadMapB f m = fmap (eithermap f) m
  where eithermap :: (b -> c) -> Either a b -> Either a c
        eithermap f (Left  a) = Left  a
        eithermap f (Right b) = Right $ f b











eitherDeNest :: Either a (Either a b) -> Either a b
eitherDeNest (Right x) = x
eitherDeNest (Left  l) = Left l











mapSnd :: (a, b) -> (b -> c) -> (a, c)
mapSnd (a, b) f = (a, f b)











-- !!Hard Hat Zone!!
--getFileIds :: String -> FilePath -> IO Either [String] [(Int, String)]
{-
getFileIds regex path = do
  filelines <- fmap (zip [1..]) $ fmap lines $ readFile path
  cmpRegex  <- P.compile P.compUTF8 P.execBlank regex
  cmpRegex' <- eitherMonadMapA (\(p, err) -> "Regex Error at " ++ (show p) ++ ": " ++ err) cmpRegex
  idlines   <- eitherMonadMapB (\regexpr  -> map (mapSnd P.execute regexpr) filelines) cmpRegex'
  return idlines
-}
